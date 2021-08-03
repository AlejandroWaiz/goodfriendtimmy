package muxadapter

import "github.com/AlejandroWaiz/goodfriendtimmy/internal/domain"

type MuxAdapter struct {
	domainport domain.DomainPort
}

func CreateMuxAdapter(domain domain.DomainPort) *MuxAdapter {
	return &MuxAdapter{domainport: domain}
}

func (ma *MuxAdapter) ListenAndServe() {

}
