class Cpdr < Formula
    desc "CLI tool to copy directories recursively"
    homepage "https://github.com/IHaveASegway/cpdr"
    version "0.0.3"
    url "https://github.com/IHaveASegway/cpdr/archive/refs/heads/main.tar.gz"
    sha256 "19a19a09c9efeb691dcffd78f3640802fb40e4c55bc8e920116842e968b77633"
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