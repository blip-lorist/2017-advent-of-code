module SpiralMemory
  class Grid
    START_POINT = [0,0]
    MOVE_RIGHT = [1,0]
    MOVE_UP = [0,1]
    MOVE_LEFT = [-1,0]
    MOVE_DOWN = [0,-1]
    DIRECTION_SEQUENCE = [MOVE_RIGHT, MOVE_UP, MOVE_LEFT, MOVE_DOWN]

    attr_accessor :squares
    def initialize
      @squares =  {coordinate_string(START_POINT) => Square.new(1, START_POINT[0], START_POINT[1])} # contains "x-y" :Square
      @last_coordinate_added = nil
      @last_move_direction = nil
    end

    def coordinate_string(point)
      # Makes a key like "1-3" out of an array [1,3]
      point.join("-")
    end

    def add_squares(value)
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
