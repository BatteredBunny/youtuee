{
  buildGoModule,
  lib,
  makeWrapper,
  yt-dlp,
}:
buildGoModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-WzP3y1DgJQ5XZmg0YwK0UJz5TiV+so01DGA5lg6DaJs=";

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
