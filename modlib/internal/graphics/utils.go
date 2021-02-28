package graphics

import "hlinspect/internal/gamelibs/hw"

func drawAACuboid(corner1, corner2 [3]float32) {

}

func drawPyramid(origin [3]float32, width, height float32) {
	halfWidth := width * 0.5
	offsets := [5][3]float32{
		[3]float32{halfWidth, halfWidth, 0},
		[3]float32{halfWidth, -halfWidth, 0},
		[3]float32{-halfWidth, -halfWidth, 0},
		[3]float32{-halfWidth, halfWidth, 0},
		[3]float32{halfWidth, halfWidth},
	}

	hw.TriGLBegin(hw.TriQuads)
	for _, offset := range offsets {
		hw.TriGLVertex3fv([3]float32{origin[0] + offset[0], origin[1] + offset[1], origin[2] + offset[2]})
	}
	hw.TriGLEnd()

	hw.TriGLBegin(hw.TriTriangleFan)
	hw.TriGLVertex3fv([3]float32{origin[0], origin[1], origin[2] + height})
	for _, offset := range offsets {
		hw.TriGLVertex3fv([3]float32{origin[0] + offset[0], origin[1] + offset[1], origin[2] + offset[2]})
	}
	hw.TriGLEnd()
}

func worldToHUDScreen(point [3]float32, width, height int) (screen [2]int, clipped bool) {
	fscreen, clipped := hw.ScreenTransform(point)
	screen[0] = int((1 + fscreen[0]) * 0.5 * float32(width))
	screen[1] = int((1 - fscreen[1]) * 0.5 * float32(height))
	return
}
