class DailyVerse < Formula
  desc "CLI tool for displaying Bible verses in the terminal"
  homepage "https://github.com/patlehmann1/daily-verse"
  url "https://github.com/patlehmann1/daily-verse/archive/refs/tags/v1.0.0.tar.gz"
  sha256 "27b057c927fc53b6e2679b09a996e1e75a145926886e9642c39497aa6398195c"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w")
  end

  test do
    # Test version output
    output = shell_output(bin/"daily-verse --version")
    assert_match "daily-verse v1.0.0", output

    # Test --daily flag (should return same verse when called twice)
    daily_output1 = shell_output(bin/"daily-verse --daily")
    assert_match(/\w+/, daily_output1)
    assert_match(/\d+:\d+/, daily_output1) # Should contain verse reference like "1:1"

    daily_output2 = shell_output(bin/"daily-verse --daily")
    assert_equal daily_output1, daily_output2, "Daily verse should be the same throughout the day"

    # Test --testament filter (old testament)
    old_output = shell_output(bin/"daily-verse --testament old")
    assert_match(/\w+/, old_output)
    assert_match(/\d+:\d+/, old_output)

    # Test --testament filter (new testament)
    new_output = shell_output(bin/"daily-verse --testament new")
    assert_match(/\w+/, new_output)
    assert_match(/\d+:\d+/, new_output)

    # Test --book filter
    psalm_output = shell_output(bin/"daily-verse --book Psalm")
    assert_match(/\w+/, psalm_output)
    assert_match(/Psalm/, psalm_output)

    # Test invalid testament input (should fail)
    output = shell_output(bin/"daily-verse --testament invalid 2>&1", 1)
    assert_match(/testament must be 'old' or 'new'/, output)

    # Test random verse (should produce valid output)
    random_output = shell_output(bin/"daily-verse")
    assert_match(/\w+/, random_output)
    assert_match(/\d+:\d+/, random_output)
  end
end
