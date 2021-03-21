package p1797

type AuthenticationManager struct {
	timeToLive int
	tokens     map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	tokens := make(map[string]int)
	return AuthenticationManager{timeToLive, tokens}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokens[tokenId] = currentTime + this.timeToLive
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if expire, found := this.tokens[tokenId]; found && expire >= currentTime {
		this.tokens[tokenId] = currentTime + this.timeToLive
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	valid := make(map[string]int)
	for k, expire := range this.tokens {
		if expire >= currentTime {
			valid[k] = expire
		}
	}
	this.tokens = valid
	return len(this.tokens)
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
