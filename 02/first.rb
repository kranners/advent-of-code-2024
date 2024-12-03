include Enumerable

lines = File.open("./input.txt", "r").readlines

def is_all_descending(record)
  record.each_cons(2) do |left, right|
    return false if left < right
  end

  return true
end

def is_all_ascending(record)
  record.each_cons(2) do |left, right|
    return false if left > right
  end

  return true
end

def is_within_range(record)
  record.each_cons(2) do |left, right|
    difference = (left - right).abs
    return false if difference < 1
    return false if difference > 3
  end

  return true
end

records = lines.map { |line|
  line.split.map { |entry| Integer(entry) }
}

puts records.map { |record| 
  (is_all_descending(record) || is_all_ascending(record)) && is_within_range(record)
}.count(true)
