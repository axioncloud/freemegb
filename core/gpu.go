package core

import (
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/gotk3/gotk3/gtk"
)

const VertexSource string = `#version 460 core
    in vec4 position;

    void main()
    {
        gl_Position = position;
    }`

const FragmentSource string = `#version 460 core
	in	vec4 inputColor;
	out vec4 outputColor;

	void main()
	{
	    outputColor = vec4(1.0, 0.0, 0.0, 1.0);
	}`

// GPUType is the structure to define what's inside a GPU
//  GPU Structure
//  ================
//  ================
type GPUType struct {
	control  byte
	scrollX  byte
	scrollY  byte
	scanline byte
	tick     byte

	pos_buffer uint32
	program    uint32
}

// GPU is the exported object used in the system
//
// GPU is exported to become a shared variable in the System object
var GPU = GPUType{
	control:  0x00,
	scrollX:  0x00,
	scrollY:  0x00,
	scanline: 0x00,
	tick:     0x00,

	pos_buffer: 0,
	program:    0,
}

func (gpu *GPUType) Init(glarea *gtk.GLArea) {
	glarea.MakeCurrent()

	// Init OpenGL
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version:", version)

	// Init Shaders
	var fragmentShader, vertexShader uint32
	fragmentShader = gl.CreateShader(gl.FRAGMENT_SHADER)
	vertexShader = gl.CreateShader(gl.VERTEX_SHADER)
	fragmentSrc, fragmentFree := gl.Strs(FragmentSource)
	vertexSrc, vertexFree := gl.Strs(VertexSource)
	gl.ShaderSource(fragmentShader, 1, fragmentSrc, nil)
	gl.ShaderSource(vertexShader, 1, vertexSrc, nil)

	gl.CompileShader(fragmentShader)
	gl.CompileShader(vertexShader)

	// Free GLSL shader strings
	fragmentFree()
	vertexFree()

	// Init Shader Program
	gpu.program = gl.CreateProgram()
	gl.AttachShader(gpu.program, vertexShader)
	gl.AttachShader(gpu.program, fragmentShader)
	gl.LinkProgram(gpu.program)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
}

func (gpu *GPUType) Run(glarea *gtk.GLArea) bool {

	err := glarea.GetError()
	if err != nil {
		panic(err)
	}

	// OpenGL is initialized and ready to go
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	var vertexData = []float32{
		-0.6, -0.6, 0.0,
		0.0, 0.6, 0.0,
		0.6, -0.6, 0.0,
		0.7, -0.1, 0.0,
		0.8, 0.1, 0.0,
		0.9, -0.1, 0.0,
	}

	// Init Buffers
	var vao, vbo uint32         // declare vao & vbo
	gl.GenVertexArrays(1, &vao) // allocate vao in vram
	gl.BindVertexArray(vao)     // bind vao to opengl

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexData), gl.Ptr(vertexData), gl.STATIC_DRAW)
	var position_loc int32 = gl.GetAttribLocation(gpu.program, gl.Str("position\000"))
	if position_loc < 0 {
		fmt.Println(position_loc)
	}
	gl.EnableVertexAttribArray(uint32(position_loc))
	gl.VertexAttribPointer(uint32(position_loc), 3, gl.FLOAT, false, 0, nil)

	gl.BindVertexArray(0)
	gl.DisableVertexAttribArray(uint32(position_loc))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(gpu.program)
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.DrawArrays(gl.TRIANGLES, 4, 3)
	gl.BindVertexArray(0)
	gl.UseProgram(0)
	return true
}

func (gpu *GPUType) Destroy(glarea *gtk.GLArea) {}
