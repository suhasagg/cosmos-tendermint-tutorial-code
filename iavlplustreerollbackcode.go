type Tree struct {
    tree     *iavl.MutableTree
    versions map[int64]map[string]string
    version  int64
}

func (t *Tree) Rollback(version int64, key string) (string, error) {
    if version > t.version {
        return "", fmt.Errorf("cannot rollback to a version greater than the current version")
    }
    if version < t.oldestVersion {
        return "", fmt.Errorf("cannot rollback to a version older than the oldest version")
    }
    t.tree.Set(version)
    t.version = version

    value, ok := t.versions[version][key]
    if !ok {
        return "", fmt.Errorf("key not found in tree version")
    }
    return value, nil
}
