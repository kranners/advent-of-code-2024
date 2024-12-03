include Enumerable

lines = File.open("./input.txt", "r").readlines

def is_all_descending(record)
  record.each_cons(2) do |left, right|
    false if left < right
  end

  true
end

def is_all_ascending(record)
  record.each_cons(2) do |left, right|
    false if left > right
  end

  true
end

def is_within_range(record)
  record.each_cons(2) do |left, right|
    difference = (left - right).abs
    false if difference < 1
    false if difference > 3
  end

  true
end

def is_record_safe(record)
  false if not is_within_range(record)

  is_all_descending(record) or is_all_ascending(record)
end

def is_record_with_tolerance_safe(record)
  true if is_record_safe(record)

  record.each_with_index { |_, index| 
    removed_record = record.dup.tap{ |i| i.delete_at(index) }

    true if is_record_safe(removed_record)
  }

  false
end

records = lines.map { |line|
  line.split.map { |entry| Integer(entry) }
}

puts records.map { |record| is_record_with_tolerance_safe(record) }.count(true)
