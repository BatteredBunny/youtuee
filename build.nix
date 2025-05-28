{ buildGoModule, pkgs, lib }:
buildGoModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-CQt0w0piC54iErtj1ro+v+aIYRHOvvrv2XNBx6sh0G0=";

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
