package selfupdate

import (
	"io"
	"strings"
)

func getPublicKeys() []io.Reader {
	return []io.Reader{
		strings.NewReader(`-----BEGIN PGP PUBLIC KEY BLOCK-----

mQENBFr9of0BCAC244urDHw6ruz2lShJ8QnVbaVx0QVyZV0pXPwN7qpoA8JCa8kh
2W7R394wvTzZxsGA1TBgechWZQazG0JD1l49lCxzXUaI9D1Z+gteCj9QoiWniYQe
GEE/8uG/apXOv2jzEZ0wirKZ1TFdzb6M8SnOLXKRzeawP+KDEM/nzVGLEO6o1Vpl
RklCPZBSuTMcqYUgx//L9lU2NSmyFPQXw1Fw+xFDVyu8tHF95qRUjTUxdycC4ZFF
kNyqwfo55SykwnhZxVOeqEe4Ks2mwWf9T2L3wWRx/mPdISnTtgN3stVC56yD0/+S
KTRka8/w5xbDgxDaBf1j4RB74zjLtHpffDHtABEBAAG0J1NwZWVkc25pdGNoIDxz
cGVlZHNuaXRjaF9hZG1pbkBzaWwub3JnPokBNwQTAQgAIQUCWv2h/QIbAwULCQgH
AgYVCAkKCwIEFgIDAQIeAQIXgAAKCRARnv7eeqSika9QB/9ZwnlfVXAjuDX+1sur
nnBeVogxI5xdbvlf5unKa/jb8ZumAGBzgmdmXaBzo//J6gScmR+p+FdK5V9qc8nm
C5HohyrXzHsG3icPXfsV+dikORRPnaTTT7wMZQY4nG6tAeZNG569cEYLdxC3oIoN
tQwBdLWr0vb1WKYw/jLuRbaz5DWCBH+SKcNIDfwrOeal7cNNdKGmR1hKRnzfkq+Y
fU+Sk7I7XxBWEzb+ljt6jhstgAOWzOVfbreo6g8viNYlX4N1V59u1b3bBoyPMDqi
SAuDmryJk3+6QXiZEPMUVBNaAngP5TKoyyq3dAREAm5lHn2GmowR2WX4JmGmocYf
MG1euQENBFr9of0BCACw2Yc+CqKO4Wh3NoI8IjmSdNbPYQ2vV7DwDyFGQdUoq4Tn
YbhcTHXlNa1VryOiHs+PpIJ2A62iaz1k34z95jJE7ZHbvl7UHFF3NnaT+iktsO2u
VSkniNzIBMBtSXpFN41k4A2vI6g+Lxe0G/VZvq92jNX6LzkyEtCjY1g2gD4R+o36
qARnBk7LfPIF7XCApWEPzlMmaRY2A87VIxrKCgpfF6Ic7eKo4IRZFw9lCNeflbbL
3DSTfDgyijhXEr7imUqhxkItow7okRq6HmIK+VWEEmG99snwwgllkweeSdCRZw3l
lo6AzbJADQXVzfGBdhrefvp3Kiw9RTqcuFQDSUMhABEBAAGJAR8EGAEIAAkFAlr9
of0CGwwACgkQEZ7+3nqkopHvfAf/VoRREGUEyysgV3Jd7qnXUNY4O6oGtr2tCeuO
nr/4cvEz7gJOII8s7Na98AXMuZPjiS+HeoILLP8LBg3sKW14MBRNpdldeiZ2EwNg
pRoJUpJWW9eEGAg/5/W9ilKpSL7m/FAYYlO7Qg007dyZrZ+RhfdcESt04nrnciBF
cFDwOpZkaZxTA8OTbs7DoMKNODe0db9ZLYuuW3Se/bqteNa6jCBTOQD9OeHD5D6o
tESbAN0gSoXwjxYFUzfHAhh6wpqXhONXdwtSny4kWPaVEPo1MEivcACaQR9hnUgf
WTLCkJU0788HLusGxHUoNQJuF5J3C3zoJk75IPCkFDZvzkPMcg==
=Hc1Z
-----END PGP PUBLIC KEY BLOCK-----
`),
	}
}
