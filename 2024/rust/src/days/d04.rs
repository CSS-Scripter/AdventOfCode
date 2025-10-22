pub fn p1(inputs: Vec<String>) -> Result<String, String> {
    let grid = Grid::new(inputs);
    let count = grid.count_xmas();
    Ok(format!("{count}").to_string())
}

pub fn p2(inputs: Vec<String>) -> Result<String, String> {
    let grid = Grid::new(inputs);
    let count = grid.count_x_shaped_mas();
    Ok(format!("{count}").to_string())
}

struct Grid {
    width: usize,
    height: usize,
    field: Vec<u8>,
}

impl Grid {
    pub fn new(field: Vec<String>) -> Grid {
        let height = field.len();
        if height == 0 {
            return Grid {
                width: 0,
                height: 0,
                field: vec![],
            };
        }

        let width = field[0].len();
        let field: Vec<u8> = field
            .iter()
            .map(|s| s.as_bytes().to_owned())
            .flatten()
            .collect();

        Grid {
            width,
            height,
            field,
        }
    }

    fn find_all(&self, c: u8) -> Vec<usize> {
        let mut indices: Vec<usize> = vec![];
        for i in 0..self.field.len() {
            if self.field[i] == c {
                indices.push(i);
            }
        }

        indices
    }

    fn flat_to_xy(&self, index: usize) -> (usize, usize) {
        let y = index / self.width;
        let x = index % self.width;
        (x, y)
    }

    fn transpose_index(&self, index: usize, dx: isize, dy: isize) -> Option<usize> {
        let (x, y) = self.flat_to_xy(index);

        // Convert to signed for offset math
        let new_x = x as isize + dx;
        let new_y = y as isize + dy;

        // Check bounds
        if new_x < 0 || new_x >= self.width as isize || new_y < 0 || new_y >= self.height as isize {
            return None; // Out of bounds
        }

        // Convert back to flat index
        Some((new_y as usize) * self.width + (new_x as usize))
    }

    pub fn count_xmas(&self) -> i32 {
        let mut count: i32 = 0;

        let flat_indices = self.find_all(b'X');
        let search: [u8; 3] = [b'M', b'A', b'S'];
        let search_matrix: [[(isize, isize); 3]; 8] = [
            // TOP LEFT
            [(-1, -1), (-2, -2), (-3, -3)],
            // TOP
            [(0, -1), (0, -2), (0, -3)],
            // TOP RIGHT
            [(1, -1), (2, -2), (3, -3)],
            // RIGHT
            [(1, 0), (2, 0), (3, 0)],
            // BOTTOM RIGHT
            [(1, 1), (2, 2), (3, 3)],
            // BOTTOM
            [(0, 1), (0, 2), (0, 3)],
            // BOTTOM LEFT
            [(-1, 1), (-2, 2), (-3, 3)],
            // LEFT
            [(-1, 0), (-2, 0), (-3, 0)],
        ];

        for fi in flat_indices {
            for search_direction in search_matrix {
                let coords: Vec<usize> = search_direction
                    .clone()
                    .iter()
                    .map(|(dx, dy)| self.transpose_index(fi, *dx, *dy))
                    .filter(|opt| opt.is_some())
                    .map(|opt| opt.unwrap())
                    .collect();
                if coords.len() != 3 {
                    // out of bounds
                    continue;
                }

                let mut is_xmas = true;
                for i in 0..3 {
                    if self.field[coords[i]] != search[i] {
                        is_xmas = false;
                        break;
                    }
                }
                if is_xmas {
                    count += 1;
                }
            }
        }

        count
    }

    pub fn count_x_shaped_mas(&self) -> i32 {
        let mut count: i32 = 0;

        let flat_indices = self.find_all(b'A');

        for fi in flat_indices {
            if self.is_x_shaped_mas(fi) {
                count += 1;
            }
        }

        count
    }

    fn is_x_shaped_mas(&self, fi: usize) -> bool {
        let transposes: [[(isize, isize); 2]; 2] = [[(-1, -1), (1, 1)], [(-1, 1), (1, -1)]];
        for tp in transposes {
            let (dx1, dy1) = tp[0];
            let (dx2, dy2) = tp[1];

            let i1 = self.transpose_index(fi, dx1, dy1);
            let i2 = self.transpose_index(fi, dx2, dy2);

            if i1.is_none() || i2.is_none() {
                return false;
            }

            let i1 = i1.unwrap();
            let i2 = i2.unwrap();

            let c1 = self.field[i1];
            let c2 = self.field[i2];

            if !((c1 == b'M' || c1 == b'S') && (c2 == b'M' || c2 == b'S') && (c1 != c2)) {
                return false;
            }
        }
        return true;
    }
}
