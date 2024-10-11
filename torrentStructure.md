

 d
  8:announce
    41:http://bttracker.debian.org:6969/announce
  7:comment
    35:"Debian CD from cdimage.debian.org"
  13:creation date
    i1573903810e
  4:info
    d
      6:length
        i351272960e
      4:name
        31:debian-10.2.0-amd64-netinst.iso
      12:piece length
        i262144e
      6:pieces
        26800:�����PS�^�� (binary blob of the hashes of each piece)
    e
e
---
- announce—the URL of the high tracker

- info—this maps to a dictionary whose keys are very dependent on whether one or more files are being shared:
  - files—a list of dictionaries each corresponding to a file (only when multiple files are being shared). Each dictionary has the following keys:
    - length—size of the file in bytes.
    - path—a list of strings corresponding to subdirectory names, the last of which is the actual file name

- length—size of the file in bytes (only when one file is being shared though)
- name—suggested filename where the file is to be saved (if one file)/suggested directory name where the files are to be saved (if multiple files)
- piece length—number of bytes per piece. This is commonly 28 KiB = 256 KiB = 262,144 B.
- pieces—a hash list, i.e., a concatenation of each piece's SHA-1 hash. As SHA-1 returns a 160-bit hash, pieces will be a string whose length is a multiple of 20 bytes. If the torrent contains multiple files, the pieces are formed by concatenating the files in the order they appear in the files dictionary (i.e., all pieces in the torrent are the full piece length except for the last piece, which may be shorter