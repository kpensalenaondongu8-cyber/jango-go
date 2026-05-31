package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ─────────────────────────────────────────────
// CHAPTER 1: FUNDAMENTALS
// The density ramp — characters ordered by visual weight
// ─────────────────────────────────────────────

// DensityRamp maps brightness (0.0–1.0) to an ASCII character.
// Lower index = lighter (less ink), higher = darker (more ink).
const densityRamp = " .,:;-=+*#%@█"

func brightnessToChar(b float64) rune {
	b = math.Max(0, math.Min(1, b))
	i := int(b * float64(len(densityRamp)-1))
	return rune(densityRamp[i])
}

// ─────────────────────────────────────────────
// CHAPTER 2: BASIC SHAPES
// ─────────────────────────────────────────────

// Box draws a simple ASCII box using line-drawing characters.
func drawBox(w, h int, title string) {
	top := "┌" + strings.Repeat("─", w-2) + "┐"
	mid := "│" + strings.Repeat(" ", w-2) + "│"
	bot := "└" + strings.Repeat("─", w-2) + "┘"

	fmt.Println(top)
	if title != "" {
		label := " " + title + " "
		pad := (w - 2 - len(label)) / 2
		row := "│" + strings.Repeat(" ", pad) + label + strings.Repeat(" ", w-2-pad-len(label)) + "│"
		fmt.Println(row)
		fmt.Println("├" + strings.Repeat("─", w-2) + "┤")
		for i := 0; i < h-4; i++ {
			fmt.Println(mid)
		}
	} else {
		for i := 0; i < h-2; i++ {
			fmt.Println(mid)
		}
	}
	fmt.Println(bot)
}

// Circle uses the parametric equation and corrects the ~2:1 aspect ratio
// of monospace fonts (characters are taller than wide).
func drawCircle(radius int) {
	aspect := 0.5 // char height ≈ 2× char width
	for y := -radius; y <= radius; y++ {
		row := ""
		for x := -radius * 2; x <= radius*2; x++ {
			// Normalise x back to a square coordinate system
			nx := float64(x) * aspect
			ny := float64(y)
			dist := math.Sqrt(nx*nx + ny*ny)
			if math.Abs(dist-float64(radius)) < 0.9 {
				row += "█"
			} else {
				row += " "
			}
		}
		fmt.Println(strings.TrimRight(row, " "))
	}
}

// Triangle draws a filled triangle.
func drawTriangle(height int) {
	for i := 1; i <= height; i++ {
		spaces := strings.Repeat(" ", height-i)
		stars := strings.Repeat("*", 2*i-1)
		fmt.Println(spaces + stars)
	}
	fmt.Println(strings.Repeat("*", 2*height-1))
}

// ─────────────────────────────────────────────
// CHAPTER 3: SHADING & LIGHTING
// Simulate a 3D sphere with a point light source
// ─────────────────────────────────────────────

// Sphere renders a shaded sphere.
// lightDir is the normalised light direction vector.
func drawSphere(radius int, lx, ly, lz float64) {
	// Normalise light direction
	llen := math.Sqrt(lx*lx + ly*ly + lz*lz)
	lx, ly, lz = lx/llen, ly/llen, lz/llen

	aspect := 0.45
	for y := -radius; y <= radius; y++ {
		row := ""
		for x := -radius * 2; x <= radius*2; x++ {
			nx := float64(x) * aspect
			ny := float64(y)
			d := nx*nx + ny*ny
			r2 := float64(radius * radius)

			if d > r2 {
				row += " "
				continue
			}

			// Surface normal at this point on the unit sphere
			nz := math.Sqrt(r2 - d)
			nlen := math.Sqrt(nx*nx + ny*ny + nz*nz)
			nx2, ny2, nz2 := nx/nlen, ny/nlen, nz/nlen

			// Lambertian shading: dot product of normal and light direction
			dot := nx2*lx + ny2*ly + nz2*lz
			brightness := math.Max(0, dot)

			// Ambient term so dark side isn't pitch black
			brightness = 0.1 + 0.9*brightness

			row += string(brightnessToChar(brightness))
		}
		fmt.Println(strings.TrimRight(row, " "))
	}
}

// ─────────────────────────────────────────────
// CHAPTER 4: IMAGE → ASCII
// Convert a 2D brightness grid to ASCII art
// ─────────────────────────────────────────────

// imageToASCII takes a 2D slice of brightness values [0,1]
// and renders it as ASCII art.
func imageToASCII(pixels [][]float64) string {
	var sb strings.Builder
	for _, row := range pixels {
		for _, v := range row {
			sb.WriteRune(brightnessToChar(v))
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

// generateGradient creates a simple left-to-right brightness gradient.
func generateGradient(w, h int) [][]float64 {
	grid := make([][]float64, h)
	for y := range grid {
		grid[y] = make([]float64, w)
		for x := range grid[y] {
			grid[y][x] = float64(x) / float64(w-1)
		}
	}
	return grid
}

// generateRadialGradient creates a radial brightness pattern.
func generateRadialGradient(w, h int) [][]float64 {
	grid := make([][]float64, h)
	cx, cy := float64(w)/2, float64(h)/2
	maxDist := math.Sqrt(cx*cx + cy*cy)
	for y := range grid {
		grid[y] = make([]float64, w)
		for x := range grid[y] {
			dx := float64(x) - cx
			dy := (float64(y) - cy) * 2 // correct aspect
			dist := math.Sqrt(dx*dx + dy*dy)
			grid[y][x] = 1 - math.Min(1, dist/maxDist)
		}
	}
	return grid
}

// generateNoise creates random noise — useful for textures.
func generateNoise(w, h int, seed int64) [][]float64 {
	rng := rand.New(rand.NewSource(seed))
	grid := make([][]float64, h)
	for y := range grid {
		grid[y] = make([]float64, w)
		for x := range grid[y] {
			grid[y][x] = rng.Float64()
		}
	}
	return grid
}

// ─────────────────────────────────────────────
// CHAPTER 5: ISOMETRIC CUBE
// Advanced: 3D projection onto 2D character grid
// ─────────────────────────────────────────────

// IsoCube draws a simple isometric cube made of line characters.
func drawIsoCube(size int) {
	// We'll assemble lines manually for clarity.
	// size = number of characters per edge
	s := size

	// Top face
	topLine := strings.Repeat("_", s*2)
	fmt.Printf("%s%s\n", strings.Repeat(" ", s), topLine)

	for i := 0; i < s; i++ {
		left := strings.Repeat(" ", s-i-1)
		right := strings.Repeat(" ", i)
		fmt.Printf("%s/%s\\%s/\n", left, strings.Repeat("_", s*2), right)
	}

	// Front and right faces
	for i := 0; i < s; i++ {
		frontFill := strings.Repeat("░", s*2)
		rightFill := strings.Repeat("▒", s)
		fmt.Printf("|%s|%s|\n", frontFill, rightFill)
	}

	// Bottom edge
	fmt.Printf("|%s|%s|\n",
		strings.Repeat("_", s*2),
		strings.Repeat("_", s))
}

// ─────────────────────────────────────────────
// CHAPTER 6: TEXT ART (Big Letters)
// Encode letters as 5×3 bit patterns
// ─────────────────────────────────────────────

// bigFont maps characters to 5-row, 3-col bitmaps.
var bigFont = map[rune][5]string{
	'A': {"###", "# #", "###", "# #", "# #"},
	'B': {"## ", "# #", "## ", "# #", "## "},
	'C': {"###", "#  ", "#  ", "#  ", "###"},
	'G': {"###", "#  ", "# #", "# #", "###"},
	'O': {"###", "# #", "# #", "# #", "###"},
	'S': {"###", "#  ", "###", "  #", "###"},
	'T': {"###", " # ", " # ", " # ", " # "},
	'E': {"###", "#  ", "##.", "#  ", "###"},
	'R': {"## ", "# #", "## ", "# #", "# #"},
	'I': {"###", " # ", " # ", " # ", "###"},
	' ': {"   ", "   ", "   ", "   ", "   "},
}

func bigText(s string) {
	rows := [5]strings.Builder{}
	for _, ch := range strings.ToUpper(s) {
		pattern, ok := bigFont[ch]
		if !ok {
			pattern = bigFont[' ']
		}
		for r, line := range pattern {
			rows[r].WriteString(strings.ReplaceAll(line, "#", "█"))
			rows[r].WriteString("  ") // letter spacing
		}
	}
	for _, row := range rows {
		fmt.Println(row.String())
	}
}

// ─────────────────────────────────────────────
// CHAPTER 7: ANIMATION
// Terminal animation using \r and ANSI escape codes
// ─────────────────────────────────────────────

// spinner animates a classic loading spinner in-place.
func spinner(label string, duration time.Duration) {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	deadline := time.Now().Add(duration)
	i := 0
	for time.Now().Before(deadline) {
		fmt.Printf("\r%s %s", frames[i%len(frames)], label)
		time.Sleep(80 * time.Millisecond)
		i++
	}
	fmt.Println("\r✓ " + label + "     ")
}

// movingWave draws an animated sine wave that scrolls horizontally.
func animateWave(frames int) {
	width := 60
	height := 10

	for f := 0; f < frames; f++ {
		// Move cursor up by height lines to overwrite
		if f > 0 {
			fmt.Printf("\033[%dA", height)
		}
		offset := float64(f) * 0.3
		for y := 0; y < height; y++ {
			row := make([]rune, width)
			for i := range row {
				row[i] = ' '
			}
			for x := 0; x < width; x++ {
				// Sine wave: amplitude = height/2, frequency varies with x
				waveY := int(float64(height)/2 + float64(height)/2.5*
					math.Sin(float64(x)*0.25+offset))
				if waveY == y {
					row[x] = '█'
				} else if waveY-1 == y {
					row[x] = '▓'
				} else if waveY-2 == y {
					row[x] = '░'
				}
			}
			fmt.Println(string(row))
		}
		time.Sleep(50 * time.Millisecond)
	}
}

// ─────────────────────────────────────────────
// CHAPTER 8: PLASMA EFFECT (mastery-level)
// Classic demoscene effect: layered sine waves
// ─────────────────────────────────────────────

const plasmaRamp = " .,:;-=+*#%@"

func plasma(w, h int, t float64) string {
	var sb strings.Builder
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Combine multiple sine waves
			v := math.Sin(float64(x)*0.3+t) +
				math.Sin(float64(y)*0.3+t) +
				math.Sin((float64(x)+float64(y))*0.2+t) +
				math.Sin(math.Sqrt(float64(x*x+y*y))*0.3)
			// Normalise from [-4,4] to [0,1]
			norm := (v + 4) / 8
			i := int(norm * float64(len(plasmaRamp)-1))
			sb.WriteByte(plasmaRamp[i])
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func animatePlasma(frames int) {
	w, h := 60, 20
	for f := 0; f < frames; f++ {
		if f > 0 {
			fmt.Printf("\033[%dA", h)
		}
		t := float64(f) * 0.15
		fmt.Print(plasma(w, h, t))
		time.Sleep(50 * time.Millisecond)
	}
}

// ─────────────────────────────────────────────
// HELPERS
// ─────────────────────────────────────────────

func header(s string) {
	line := strings.Repeat("═", len(s)+4)
	fmt.Printf("\n╔%s╗\n║  %s  ║\n╚%s╝\n\n", line, s, line)
}

func pause() {
	fmt.Println("\n  ── press Enter to continue ──")
	fmt.Scanln()
}

// ─────────────────────────────────────────────
// MAIN — interactive tour
// ─────────────────────────────────────────────

func main() {
	interactive := len(os.Args) == 1 // pause between chapters only in interactive mode

	// ── Banner ──────────────────────────────
	fmt.Println()
	bigText("ASCII ART")
	fmt.Println()
	fmt.Println("  Fundamentals to Mastery · Go Edition")
	fmt.Println()

	if interactive {
		pause()
	}

	// ── Chapter 1: Basic shapes ─────────────
	header("Chapter 1 · Basic Shapes")

	fmt.Println("── Box with title ──")
	drawBox(30, 5, "Hello, ASCII")
	fmt.Println()

	fmt.Println("── Circle (aspect-ratio corrected) ──")
	drawCircle(6)
	fmt.Println()

	fmt.Println("── Triangle ──")
	drawTriangle(7)
	fmt.Println()

	if interactive {
		pause()
	}

	// ── Chapter 2: Gradients ────────────────
	header("Chapter 2 · Gradients & Ramps")

	fmt.Println("── Density ramp ──")
	fmt.Printf("  %s\n\n", densityRamp)

	fmt.Println("── Linear gradient (left → right) ──")
	grid := generateGradient(60, 5)
	fmt.Print(imageToASCII(grid))

	fmt.Println("── Radial gradient ──")
	radial := generateRadialGradient(60, 15)
	fmt.Print(imageToASCII(radial))

	fmt.Println("── Noise texture ──")
	noise := generateNoise(60, 10, 42)
	fmt.Print(imageToASCII(noise))

	if interactive {
		pause()
	}

	// ── Chapter 3: 3D Sphere ────────────────
	header("Chapter 3 · Shaded Sphere (Lambertian lighting)")

	fmt.Println("── Light from top-left ──")
	drawSphere(10, -1, -1, 1)
	fmt.Println()

	fmt.Println("── Light from front ──")
	drawSphere(10, 0, 0, 1)
	fmt.Println()

	if interactive {
		pause()
	}

	// ── Chapter 4: Isometric Cube ───────────
	header("Chapter 4 · Isometric Cube")
	drawIsoCube(4)
	fmt.Println()

	if interactive {
		pause()
	}

	// ── Chapter 5: Big Text ─────────────────
	header("Chapter 5 · Big Text (5×3 bitmap font)")
	bigText("GO ART")
	fmt.Println()

	if interactive {
		pause()
	}

	// ── Chapter 6: Animation ────────────────
	header("Chapter 6 · Animation")

	fmt.Println("── Spinner ──")
	spinner("Loading assets", 2*time.Second)
	spinner("Compiling shaders", 1500*time.Millisecond)

	fmt.Println("\n── Animated sine wave ──")
	animateWave(60)
	fmt.Println()

	if interactive {
		pause()
	}

	// ── Chapter 7: Plasma (Mastery) ─────────
	header("Chapter 7 · Plasma Effect (Mastery)")
	fmt.Println("Running for 3 seconds…")
	animatePlasma(60)
	fmt.Println()

	// ── Done ────────────────────────────────
	header("Complete!")
	fmt.Println("  Key concepts covered:")
	fmt.Println("  · Density ramps & brightness mapping")
	fmt.Println("  · Aspect-ratio correction for circles")
	fmt.Println("  · Lambertian (diffuse) shading on a sphere")
	fmt.Println("  · Isometric projection")
	fmt.Println("  · Bitmap font rendering")
	fmt.Println("  · Terminal animation with ANSI escape codes")
	fmt.Println("  · Demoscene plasma effect via layered sine waves")
	fmt.Println()
	fmt.Println("  Run with any argument to skip pauses:")
	fmt.Println("  go run ascii_art.go auto")
	fmt.Println()
}
