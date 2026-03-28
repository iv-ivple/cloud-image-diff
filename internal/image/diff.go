package image

type DiffResult struct {
    OnlyInA  []ImageMeta
    OnlyInB  []ImageMeta
    Changed  []ImageChange
    Matching []ImageMeta
}

type ImageChange struct {
    A, B   ImageMeta
    Fields []string
}

func imageKey(img ImageMeta) string {
    return img.Release + "|" + img.Arch + "|" + img.Region
}

func Diff(a, b []ImageMeta) DiffResult {
    mapA := make(map[string]ImageMeta, len(a))
    mapB := make(map[string]ImageMeta, len(b))
    for _, img := range a { mapA[imageKey(img)] = img }
    for _, img := range b { mapB[imageKey(img)] = img }

    result := DiffResult{}
    for key, imgA := range mapA {
        imgB, found := mapB[key]
        if !found {
            result.OnlyInA = append(result.OnlyInA, imgA)
            continue
        }
        if changed := diffFields(imgA, imgB); len(changed) > 0 {
            result.Changed = append(result.Changed, ImageChange{imgA, imgB, changed})
        } else {
            result.Matching = append(result.Matching, imgA)
        }
    }
    for key, imgB := range mapB {
        if _, found := mapA[key]; !found {
            result.OnlyInB = append(result.OnlyInB, imgB)
        }
    }
    return result
}
