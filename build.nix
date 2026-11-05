{
  buildGoLatestModule,
  lib,
  makeWrapper,
  yt-dlp,
}:
buildGoLatestModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-tt8yHm56/LEc88/Gbf6ElQdm8hfFJwZcyr0sBdHO5R0=";

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
