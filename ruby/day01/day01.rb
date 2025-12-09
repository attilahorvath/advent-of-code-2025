#! /usr/bin/env ruby

class Dial
  attr_reader :position, :zeros

  def initialize
    @position = 50
    @zeros = 0
  end

  def rotate(direction, distance)
    if direction == 'L'
      @position -= distance
    else
      @position += distance
    end

    @position %= 100
    @zeros += 1 if @position == 0
  end

  def rotate_clicks(direction, distance)
    distance.times do
      rotate(direction, 1)
    end
  end
end

def rotate_dial(clicks: false)
  dial = Dial.new

  File.open('input.txt').each_line do |line|
    direction = line[0]
    distance = line[1..].to_i

    if clicks
      dial.rotate_clicks(direction, distance)
    else
      dial.rotate(direction, distance)
    end
  end

  dial.zeros
end

puts "zeros: #{rotate_dial}"
puts "zeros with clicks: #{rotate_dial(clicks: true)}"
