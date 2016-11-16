def func1
    puts "sleep 3 seconds in func1\n"
      sleep(3)
    Time.now
end

threads = []
10000.times do |i|
threads << Thread.new { func1  }
end

threads.each { |t|puts t.join  }
