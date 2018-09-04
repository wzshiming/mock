package mock

import (
	"github.com/wzshiming/crun"
)

// RandDomain Returns a random domain
func RandDomain() string {
	return domain.Rand()
}

// RandURL Returns a random URL
func RandURL() string {
	return url.Rand()
}

// RandUUID Returns a random UUID
func RandUUID() string {
	return uuid.Rand()
}

// RandEmail Returns a random Email
func RandEmail() string {
	return email.Rand()
}

// RandName Returns a random name
func RandName() string {
	return name.Rand()
}

// crun constant
var (
	protocol = crun.MustCompile(_protocol)
	domain   = crun.MustCompile(_domain)
	url      = crun.MustCompile(_url)
	email    = crun.MustCompile(_email)
	uuid     = crun.MustCompile(_uuid)
	name     = crun.MustCompile(_name)
)

const (
	_firstName = `(James|John|Robert|Michael|William|David|Richard|Charles|Joseph|Thomas|Christopher|Daniel|Paul|Mark|Donald|George|Kenneth|Steven|Edward|Brian|Ronald|Anthony|Kevin|Jason|Matthew|Gary|Timothy|Jose|Larry|Jeffrey|Frank|Scott|Eric|Mary|Patricia|Linda|Barbara|Elizabeth|Jennifer|Maria|Susan|Margaret|Dorothy|Lisa|Nancy|Karen|Betty|Helen|Sandra|Donna|Carol|Ruth|Sharon|Michelle|Laura|Sarah|Kimberly|Deborah|Jessica|Shirley|Cynthia|Angela|Melissa|Brenda|Amy|Anna)`
	_lastName  = `(Smith|Johnson|Williams|Brown|Jones|Miller|Davis|Garcia|Rodriguez|Wilson|Martinez|Anderson|Taylor|Thomas|Hernandez|Moore|Martin|Jackson|Thompson|White|Lopez|Lee|Gonzalez|Harris|Clark|Lewis|Robinson|Walker|Perez|Hall|Young|Allen)`
	_name      = `((` + _firstName + ` )?` + _lastName + `)`
	_topDomain = `(com|net|org|edu|gov|int|mil|tel|biz|cc|tv|info|name|hk|mobi|asia|cd|travel|pro|museum|coop|aero|ad|ae|af|ag|ai|al|am|an|ao|aq|ar|as|at|au|aw|az|ba|bb|bd|be|bf|bg|bh|bi|bj|bm|bn|bo|br|bs|bt|bv|bw|by|bz|ca|cc|cf|cg|ch|ci|ck|cl|cm|cn|co|cq|cr|cu|cv|cx|cy|cz|de|dj|dk|dm|do|dz|ec|ee|eg|eh|es|et|ev|fi|fj|fk|fm|fo|fr|ga|gb|gd|ge|gf|gh|gi|gl|gm|gn|gp|gr|gt|gu|gw|gy|hk|hm|hn|hr|ht|hu|id|ie|il|in|io|iq|ir|is|it|jm|jo|jp|ke|kg|kh|ki|km|kn|kp|kr|kw|ky|kz|la|lb|lc|li|lk|lr|ls|lt|lu|lv|ly|ma|mc|md|mg|mh|ml|mm|mn|mo|mp|mq|mr|ms|mt|mv|mw|mx|my|mz|na|nc|ne|nf|ng|ni|nl|no|np|nr|nt|nu|nz|om|qa|pa|pe|pf|pg|ph|pk|pl|pm|pn|pr|pt|pw|py|re|ro|ru|rw|sa|sb|sc|sd|se|sg|sh|si|sj|sk|sl|sm|sn|so|sr|st|su|sy|sz|tc|td|tf|tg|th|tj|tk|tm|tn|to|tp|tr|tt|tv|tw|tz|ua|ug|uk|us|uy|va|vc|ve|vg|vn|vu|wf|ws|ye|yu|za|zm|zr|zw)`
	_protocol  = `(ws|wss|http|https|ftp|gopher|mailto|mid|cid|news|nntp|prospero|telnet|rlogin|tn3270|wais|unix)`
	_domain    = `(` + _word + `\.` + _topDomain + `)`
	_pair      = `(` + _word + `=` + _randData + `)`
	_args      = `(` + _pair + `(&` + _pair + `){0,5})`
	_path      = `((/` + _randData + `){1,4}/?)`
	_url       = `(` + _protocol + `://` + _domain + `(` + _path + `(\?` + _args + `)?)?)`
	_email     = `(` + _randData + `@` + _domain + `)`
	_hex       = `([0-9a-f])`
	_word      = `([a-z]{2,10}|` + _protocol + `|` + _topDomain + `)`
	_number    = `([1-9][0-9]{0,8})`
	_uuid      = `(` + _hex + `{8}(-` + _hex + `{4}){3}-` + _hex + `{12})`
	_randData  = `(` + _word + `|` + _number + `)`
)
