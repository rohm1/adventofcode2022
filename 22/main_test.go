package main

import "testing"

func TestFaceChange1(t *testing.T) {
	np, nd := determineNextPos2(Down, pos{left: 109, top: 49})
	if np.left != 99 || np.top != 59 {
		t.Errorf("wrong new position: left: %d, top: %d", np.left, np.top)
	}

	if nd != Left {
		t.Errorf("wrong new direction: %d", nd)
	}
}

func TestFaceChange2(t *testing.T) {
	np, nd := determineNextPos2(Down, pos{left: 66, top: 149})
	if np.left != 49 || np.top != 166 {
		t.Errorf("wrong new position: left: %d, top: %d", np.left, np.top)
	}

	if nd != Left {
		t.Errorf("wrong new direction: %d", nd)
	}
}

func TestFaceChange3(t *testing.T) {
	np, nd := determineNextPos2(Right, pos{left: 49, top: 160})
	if np.left != 60 || np.top != 149 {
		t.Errorf("wrong new position: left: %d, top: %d", np.left, np.top)
	}

	if nd != Up {
		t.Errorf("wrong new direction: %d", nd)
	}
}

func TestFaceChange4(t *testing.T) {
	np, nd := determineNextPos2(Right, pos{left: 149, top: 47})
	if np.left != 99 || np.top != 102 {
		t.Errorf("wrong new position: left: %d, top: %d", np.left, np.top)
	}

	if nd != Left {
		t.Errorf("wrong new direction: %d", nd)
	}
}

func TestFaceChange5(t *testing.T) {
	np, nd := determineNextPos2(Right, pos{left: 99, top: 47})
	if np.left != 100 || np.top != 47 {
		t.Errorf("wrong new position: left: %d, top: %d", np.left, np.top)
	}

	if nd != Right {
		t.Errorf("wrong new direction: %d", nd)
	}
}

func TestSameFace1(t *testing.T) {
	np, nd := determineNextPos2(Right, pos{left: 50, top: 0})
	if np.left != 51 || np.top != 0 {
		t.Errorf("wrong new position: left: %d, top: %d", np.left, np.top)
	}

	if nd != Right {
		t.Errorf("wrong new direction: %d", nd)
	}
}
