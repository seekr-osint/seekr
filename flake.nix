{
  description = "A simple Go package";

  # Nixpkgs / NixOS version to use.
  inputs = {
    nixpkgs = {
      url = "github:Nixos/nixpkgs/nixpkgs-unstable";
      flake = true;
    };
    flake-compat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, flake-compat }:
    let

      # to work with older version of flakes
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";

      # Generate a user-friendly version number.
      version = builtins.substring 0 8 lastModifiedDate;

      # System types to support.
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];

      # Helper function to generate an attrset '{ x86_64-linux = f "x86_64-linux"; ... }'.
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

      # Nixpkgs instantiated for supported system types.
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });

    in
    {

      # Provide some binary packages for selected system types.
      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          seekr = pkgs.buildGoModule {
            pname = "seekr";
            inherit version;
            # In 'nix develop', we don't need a copy of the source tree
            # in the Nix store.
            src = ./.;

            #vendorSha256 = pkgs.lib.fakeSha256;
            vendorSha256 = "sha256-Alq3l3/bpDCqK8d0wkgqoa0Jyi+J9linXgCiQD6QRt4=";
          };
        });

      apps = forAllSystems (system: {
        default = {
          type = "app";
          program = "${self.packages.${system}.seekr}/bin/seekr";
        };
      });

      formatter = forAllSystems (system: nixpkgsFor.${system}.nixpkgs-fmt);

      devShells = forAllSystems (system: {
        default = nixpkgsFor.${system}.mkShell {
          packages = [
            nixpkgsFor.${system}.go
            # jq is useful to debug the database
            nixpkgsFor.${system}.jq
            nixpkgsFor.${system}.maigret
            nixpkgsFor.${system}.goreleaser
            self.packages.${system}.seekr
          ];
        };
      });


      defaultPackage = forAllSystems (system: self.packages.${system}.seekr);
    };
}
