#! /usr/bin/env ruby

class Database
  def initialize
    @ranges = []
    @ids = []
  end

  def add_range(min, max)
    @ranges.push(min..max)
  end

  def add_id(id)
    @ids.push(id)
  end

  def join_ranges
    raw_ranges = @ranges.sort_by(&:begin)
    @ranges.clear

    min = nil
    max = nil

    raw_ranges.each do |range|
      unless min
        min = range.begin
        max = range.end

        next
      end

      if range.begin > max
        add_range(min, max)

        min = range.begin
        max = range.end
      else
        max = [max, range.end].max
      end
    end

    add_range(min, max) if min && max
  end

  def fresh_ids
    @ids.select { |id| @ranges.any? { |range| range.cover?(id) } }.count
  end

  def total_fresh_ids_in_ranges
    @ranges.map { |range| range.end - range.begin + 1 }.sum
  end
end

def load_database
  database = Database.new
  ranges = true

  File.open('input.txt').each_line do |line|
    if ranges
      if line == "\n"
        ranges = false
        next
      end

      min, max = line.split('-').map(&:to_i)
      database.add_range(min, max)
    else
      database.add_id(line.to_i)
    end
  end

  database.join_ranges

  database
end

database = load_database

puts "fresh ids: #{database.fresh_ids}"
puts "total fresh ids in ranges: #{database.total_fresh_ids_in_ranges}"
