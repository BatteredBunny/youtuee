{
  buildGoModule,
  lib,
  makeWrapper,
  yt-dlp,
}:
buildGoModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-hX8lq6UAd4s3Z52YhdsaMwZ6M/cHZyM151UwZttHz+c=";

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
