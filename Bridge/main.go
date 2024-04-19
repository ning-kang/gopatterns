package main

import "fmt"

type Compiler interface {
	CompileGo(app GoApp)
	CompilePython(app PythonApp)
}

type WindowsCompiler struct {
	OSVersion string
}

func (w *WindowsCompiler) CompileGo(app GoApp) {
	fmt.Printf("Compiling %s on %s: %s\n", app.name, w.OSVersion, app.code)
}

func (w *WindowsCompiler) CompilePython(app PythonApp) {
	fmt.Printf("Compiling %s on %s: %s\n", app.name, w.OSVersion, app.code)
}

type LinuxCompiler struct {
	OSVersion string
}

func (l *LinuxCompiler) CompileGo(app GoApp) {
	fmt.Printf("Compiling %s on %s: %s\n", app.name, l.OSVersion, app.code)
}

func (l *LinuxCompiler) CompilePython(app PythonApp) {
	fmt.Printf("Compiling %s on %s: %s\n", app.name, l.OSVersion, app.code)
}

type PythonApp struct {
	compiler Compiler // This allows both WindowsCompiler and LinuxCompiler to be passed in because they both satisfy Compiler interface
	name     string
	code     string
}

func NewPythonApp(compiler Compiler, name string, code string) *PythonApp {
	return &PythonApp{compiler: compiler, name: name, code: code}
}

func (p *PythonApp) Compile() {
	p.compiler.CompilePython(*p)
}

type GoApp struct {
	compiler Compiler
	name     string
	code     string
}

func NewGoApp(compiler Compiler, name string, code string) *GoApp {
	return &GoApp{compiler: compiler, name: name, code: code}
}

func (g *GoApp) Compile() {
	g.compiler.CompileGo(*g)
}

func main() {
	w := &WindowsCompiler{OSVersion: "Windows 11"}
	l := &LinuxCompiler{OSVersion: "Ubuntu 20.04"}

	goapp := NewGoApp(w, "my go app", "fmt.Print()")
	pythonapp := NewPythonApp(l, "my python app", "print()")

	goapp.Compile()
	pythonapp.Compile()
}
