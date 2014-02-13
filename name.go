package spki

// A Name represents local & extended SPKI names, as well as simple
// principals which are just a key.  A local name will have one name
// in Names; an extenedd name will have multiple names.  A simple
// principal will have Principal but no Names.
type Name struct {
	Principal Key
	Names []string
}

// IsPrincipal returns true if n is a principal name, i.e. if it
// refers directly to a key and no names in that key's namespace.
func (n *Name) IsPrincipal() bool {
	if n.Principal != nil && len(n.Names) == 0 {
		return true
	}
	return false
}

// Local returns the local part of n, e.g. (name #123# a b c) would
// return (name #123# a).
func (n *Name) Local() *Name {
	if len(n.Names) < 2 {
		return n
	}
	return &Name{n.Principal, n.Names[0:1]}
}

	// IsLocal returns true if n is a local name—i.e., len(n.Names) is 0 or 1
func (n *Name) IsLocal() bool {
	return len(n.Names) < 2
}

// IsPrefix returns true if n is a prefix of n2
func (n *Name) IsPrefix(n2 *Name) bool {
	if n == nil {
		return n2 == nil
	}
	if n.Principal != nil && !n.Principal.Equal(n2.Principal) {
		return false
	}
	for names1, names2, i := n.Names, n2.Names, 0; i < len(names1) && i < len(names2); i++ {
		if names1[i] != names2[i] {
			return false
		}
	}
	return true
}