class Cpdr < Formula
    desc "CLI tool to copy directories recursively"
    homepage "https://github.com/IHaveASegway/cpdr"
    version "1.0.3"
    url "https://github.com/ihaveasegway/cpdr/releases/download/v1.0.3/cpdr-1.0.3.tar.gz" # updated URL
    sha256 "dd2d1e2dd08d9e05ec4823ca1ea9fe3030affbfec838e64ca7948d4965620d3b" # Update with the actual SHA256 hash for v1.0.3
    license "MIT"
  
    depends_on "go" => :build
  
    def install
      system "go", "build", "-o", "cpdr", "cpdr.go"
      bin.install "cpdr"
    end
  
    test do
      assert_match "cpdr", shell_output("#{bin}/cpdr --help 2>&1", 1)
    end
end