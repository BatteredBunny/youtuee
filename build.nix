{ buildGoModule, pkgs, lib }:
buildGoModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-YwSo2TJqJdLfyjAG+zlHD+hdYbUqk/fyrXd5xBHTPRM=";

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
