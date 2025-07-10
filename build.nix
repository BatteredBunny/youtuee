{ buildGoModule, pkgs, lib }:
buildGoModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-Bj1OUYz9XjQxoKyP3SpkC7fHtT7M0bcYE3NPyDX0+4E=";

  ldflags = [
    "-s"
    "-w"
  ];

  nativeBuildInputs = with pkgs; [
    makeWrapper
  ];

  postInstall = ''
    wrapProgram $out/bin/youtuee --prefix PATH : ${lib.makeBinPath [ pkgs.yt-dlp ]}
  '';

  meta = {
    homepage = "https://github.com/BatteredBunny/youtuee";
    mainProgram = "youtuee";
  };
}
