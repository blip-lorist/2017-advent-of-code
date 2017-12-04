require 'minitest/autorun'
require_relative 'spiral_memory'

describe SpiralMemory::Grid do
  before do
    @grid = SpiralMemory::Grid.new
  end

  it "has a squares attribute" do
    assert(@grid.squares.is_a?(Hash))
  end

  it "places the first square" do
    assert(@grid.squares["0-0"].is_a?(SpiralMemory::Square))
  end

  describe "#get_distance" do
    it "tells us the taxicab distance from square N to 1" do
      assert_equal(0, @grid.get_distance(1))
    end

    it "places the next square" do
     assert_equal(1, @grid.get_distance(2))
     assert_equal(@grid.last_square.value, 1) # part two adjacent square value
     assert_equal(@grid.last_square.row, 0)
     assert_equal(@grid.last_square.col, 1)
    end

    it "handles repetitive directions" do
      assert_equal(2, @grid.get_distance(5))
    end

    it "moves further out" do
      assert_equal(31, @grid.get_distance(1024))
    end
  end

  describe "#squareBuilder" do
    it "adds a square to the grid" do
      @grid.add_squares(2)
    end
  end

  describe "#find_adjacent_values" do
    it "finds all adjacent values" do
      @grid.add_squares(3) # this is placed at 1-1
      third_square = @grid.squares["1-1"]
      assert_equal([1, 1], @grid.find_adjacent_values(third_square))
    end
  end

  describe "#square values attribute" do
    it "is set to the sum of adjacent square values" do
      @grid.add_squares(3) # this is placed at 1-1
      assert_equal(1, @grid.last_square.col)
      assert_equal(1, @grid.last_square.row)
      third_square = @grid.last_square
      assert_equal(2, third_square.value)
    end

    it "still works on the 10th square..." do
      @grid.add_squares(10)
      tenth_square  = @grid.last_square
      assert_equal(26, tenth_square.value)
    end
  end
end
