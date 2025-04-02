class Cpdr < Formula
    desc "A cool Go script"
    homepage "https://github.com/ihaveasegway/cpdr"
    url "https://github.com/ihaveasegway/cpdr/releases/download/v1.0.2/cpdr-1.0.2.tar.gz"
    sha256 "<SHA256_HASH>" # Replace with the hash
    version "1.0.2"
  
    depends_on "go" => :build
  
    def install
      system "go", "build", "-o", "cpdr", "cpdr.go"
      bin.install "cpdr"
    end
  
    test do
      assert_match "expected output", shell_output("#{bin}/cpdr") # Adjust to 1.0.2’s output
    end
  end