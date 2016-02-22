package rpc

func (s *Sender) Sign() []byte {
	var app [16]byte
	var tag [36]byte

	copy(app[:], s.name)
	copy(tag[:], s.uuid)

	var body []byte

	body = append(body[:], app[:]...)
	body = append(body[:], tag[:]...)

	return body
}

func (p *Receiver) Sign () []byte {

	var body []byte

	var name [16]byte
	var uuid [36]byte
	var hander [16]byte

	copy(name[:], p.name)
	copy(uuid[:], p.uuid)
	copy(hander[:], p.handler)

	body = append(body, name[:]...)
	body = append(body, uuid[:]...)
	body = append(body, hander[:]...)

	return body
}

func (r *Destination) Sign () []byte {

	var body []byte

	var name [16]byte
	var uuid [36]byte
	var hander [16]byte

	copy(name[:], r.name)
	copy(uuid[:], r.uuid)
	copy(hander[:], r.handler)

	body = append(body, name[:]...)
	body = append(body, uuid[:]...)
	body = append(body, hander[:]...)

	return body
}

func (r *RPC) encode ( s Sender, d Destination, p Receiver, data []byte) []byte {
	var body []byte
	var hash [256]byte

	var token [32]byte
	copy(token[:], r.token)

	copy(hash[0:32], token[:])
	copy(hash[32:84], s.Sign()[:])
	copy(hash[84:152], d.Sign()[:])
	copy(hash[152:220], p.Sign()[:])


	body = append(body, hash[:]...)
	body = append(body, data[:]...)

	return body
}