module SpiralMemory
  class Grid
    START_POINT = [0,0]
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

    def coordinate_string(points)
      # Makes a key like "1-3" out of an array [1,3]
      points.join("-")
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
      end

      # Add the new square to our grid
      self.squares[coordinate_string(new_coordinates)] = Square.new(value, new_coordinates[0], new_coordinates[1])

      self.last_coordinate_added = coordinate_string(new_coordinates)
      self.last_move_direction = move_direction
    end

    def get_new_coordinates(direction)
      new_row = self.last_square.row + direction[0]
      new_col = self.last_square.col + direction[1]
      [new_row, new_col]
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
  end

  class Square
    attr_reader :value, :row, :col
    def initialize(value, row, col)
      @value = value
      @row = row
      @col = col
    end
  end
end
