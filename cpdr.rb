class Cpdr < Formula
  desc "Tool to copy directory contents to clipboard"
  homepage "https://github.com/yourusername/cpdr"
  version "0.1.0"
  
  # For local installation, use the local path
  url "file:///Users/joseph.peterson/Documents/GitHub/cpdr"
  
  # If it depends on Go to run (not just compile)
  depends_on "go" => :build
  
  def install
    # Install Go dependencies first
    system "go", "get", "github.com/atotto/clipboard"
    system "go", "build", "-o", "cpdr"
    bin.install "cpdr"
  end
  
  test do
    system "#{bin}/cpdr", "--help"
  end
end