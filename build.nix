{
  buildGoModule,
  lib,
  makeWrapper,
  yt-dlp,
}:
buildGoModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-TM7K65j+99bR+2d+LpQbwlqEb0x/8Prm153KsjxEpT8=";

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
