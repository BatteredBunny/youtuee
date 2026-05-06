{
  buildGoModule,
  lib,
  makeWrapper,
  yt-dlp,
}:
buildGoModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-jNcLS1uW5nSiOcjUs6l33wDp2ztzdCLnBTie4Qv25KY=";

  ldflags = [
    "-s"
    "-w"
  ];

  nativeBuildInputs = [
    makeWrapper
  ];

  postInstall = ''
    wrapProgram $out/bin/youtuee --prefix PATH : ${lib.makeBinPath [ yt-dlp ]}
  '';

  meta = {
    homepage = "https://github.com/BatteredBunny/youtuee";
    mainProgram = "youtuee";
  };
}
