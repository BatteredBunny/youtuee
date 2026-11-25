{
  buildGoLatestModule,
  lib,
  makeWrapper,
  yt-dlp,
}:
buildGoLatestModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-5XbJlsQ72tjTQOPCl9XSlcMaO35d0hVQEPXkLOs7ljg=";

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
