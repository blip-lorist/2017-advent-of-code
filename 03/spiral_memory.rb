module SpiralMemory
  class Grid
    START_POINT = [0,0] #[col, row]
    MOVE_RIGHT = [1,0]
    MOVE_UP = [0,1]
    MOVE_LEFT = [-1,0]
    MOVE_DOWN = [0,-1]
    DIRECTION_SEQUENCE = [MOVE_RIGHT, MOVE_UP, MOVE_LEFT, MOVE_DOWN]

    attr_accessor :squares, :last_coordinate_added, :last_move_direction
    def initialize
      @squares =  {coordinate_string(START_POINT) => Square.new(1, START_POINT[0], START_POINT[1])} # contains "x-y" :Square
      @last_coordinate_added = coordinate_string(START_POINT)
      @last_move_direction = nil
    end

    def add_squares(greatest_value)
      (2..greatest_value).each do |value|
        place_square(value)
      end
    end

    def place_square(value)
      if self.last_move_direction == nil
        # This is the first new square, go right
        move_direction = MOVE_RIGHT
        new_coordinates = get_new_coordinates(move_direction)
      else
        # Try next direction
        move_direction = get_next_direction
        new_coordinates = get_new_coordinates(move_direction)

        if square_taken?(new_coordinates)
          # Go in previous direction
          move_direction = last_move_direction
          new_coordinates = get_new_coordinates(move_direction)
        end
      end

      # Add the new square to our grid
      new_square = Square.new(value, new_coordinates[0], new_coordinates[1])
      adjacent_summed_value = sum_adjacent_values(new_square)

      puts(adjacent_summed_value)
      # Hard-coded halt
      if adjacent_summed_value > 312051
        puts(adjacent_summed_value)
        exit
      end

      new_square.value = adjacent_summed_value

      self.squares[coordinate_string(new_coordinates)] = new_square
      self.last_coordinate_added = coordinate_string(new_coordinates)
      self.last_move_direction = move_direction
    end

    def coordinate_string(points)
      # Makes a key like "col-row" out of an array [col,row]
      points.join("-")
    end

    def square_taken?(coordinates)
      self.squares[coordinate_string(coordinates)].is_a?(Square)
    end

    def get_next_direction
      last_move_direction_index = DIRECTION_SEQUENCE.index(self.last_move_direction)
      if last_move_direction_index == (DIRECTION_SEQUENCE.length - 1)
        # Return to first direction
        next_move_direction_index = 0
      else
        next_move_direction_index = last_move_direction_index + 1
      end

      DIRECTION_SEQUENCE[next_move_direction_index]
    end
    def get_new_coordinates(direction)
      new_col = self.last_square.col + direction[0]
      new_row = self.last_square.row + direction[1]

      [new_col, new_row]
    end

    def get_distance(value)
      add_squares(value) unless value == 1
      col = self.last_square.col
      row = self.last_square.row
      col.abs + row.abs
    end

    def last_square
      self.squares[self.last_coordinate_added]
    end

    def find_adjacent_values(square)
      adjacent_cols = ((square.col - 1)..(square.col + 1))
      adjacent_rows = ((square.row - 1)..(square.row + 1))
      current_position = square.col.to_s + "-" + square.row.to_s

      adjacent_values = []
      # Check each position around a square for 
      adjacent_cols.each do |col|
        adjacent_rows.each do |row|
          col_row_combo = coordinate_string([col, row])
          unless col_row_combo == current_position
            adjacent_square = self.squares[col_row_combo]
            if adjacent_square
              adjacent_values << adjacent_square.value
            end
          end
        end
      end
      adjacent_values
    end

    def sum_adjacent_values(square)
      adjacent_values = find_adjacent_values(square)
      sum = 0
      adjacent_values.each do |value|
        sum += value
      end

      sum
    end
  end

  class Square
    attr_reader :row, :col
    attr_accessor :value
    def initialize(value, col, row)
      @value = value
      @row = row
      @col = col
    end
  end
end


# Answer part one
#puts(SpiralMemory::Grid.new.get_distance(312051))

# Answer part two
@grid = SpiralMemory::Grid.new
nth_square = 2
while true
  # Just keep adding squares til program exits

  @grid.place_square(nth_square)
  nth_square += 1
end
