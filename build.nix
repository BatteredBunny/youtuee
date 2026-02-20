{
  buildGoModule,
  lib,
  makeWrapper,
  yt-dlp,
}:
buildGoModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-ceqxvF7eh0cZvvnYu2JJCrjVpqU54liJ5lRn4mGoi5A=";

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
