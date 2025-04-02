class Cpdr < Formula
    desc "CLI tool to copy directories recursively"
    homepage "https://github.com/ihaveasegway/cpdr"
    url "https://github.com/ihaveasegway/cpdr/releases/download/v1.0.4/cpdr-1.0.4.tar.gz"
    sha256 "fed790b56a8961041534f3ff7d7e8e2e948ebfff592b386e317e2ad13932736b"
    license "MIT"
    version "1.0.4"
  
    depends_on "go" => :build
  
    def install
      system "go", "build", "-o", "cpdr", "cpdr.go"
      bin.install "cpdr"
    end
  
    test do
      assert_match "cpdr", shell_output("#{bin}/cpdr --help 2>&1", 1)
    end
  end