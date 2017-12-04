require 'minitest/autorun'
require_relative 'spiral_memory'

describe SpiralMemory::Grid do
  before do
    @grid = SpiralMemory::Grid.new
  end

  it "has a squares attribute" do
    assert(@grid.squares.is_a?(Hash))
  end

  it "has the first square" do
    assert(@grid.squares["0-0"].is_a?(SpiralMemory::Square))
  end

  #describe "#get_distance" do
  #  it "tells us the taxicab distance from square N to 1" do
  #    assert_equal(0, @grid.get_distance(1))
  #    assert_equal(3, @grid.get_distance(12))
  #    assert_equal(2, @grid.get_distance(23))
  #    assert_equal(31, @grid.get_distance(1024))
  #  end
  #end

  describe "#squareBuilder" do
    it "adds a square to the grid" do
      @grid.add_squares(2)
    end
  end
end
