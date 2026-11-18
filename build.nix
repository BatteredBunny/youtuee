{
  buildGoLatestModule,
  lib,
  makeWrapper,
  yt-dlp,
}:
buildGoLatestModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-IQjn+3u7UkoEjW6c/nom1oYQBErynOIl5YE0VhSUmPk=";

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
