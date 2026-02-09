{lib, ...}: {
  enableAll = list: conf:
    lib.foldl' lib.recursiveUpdate (lib.listToAttrs (lib.map (name: {
        inherit name;
        value = {enable = true;};
      })
      list))
    [conf];
}
