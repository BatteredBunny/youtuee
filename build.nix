{
  buildGoLatestModule,
  lib,
  makeWrapper,
  yt-dlp,
}:
buildGoLatestModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-yeXN1ti1yEb9+hHApvoKrC9Wf8Mf3LfEn1SwXX69XVE=";

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
