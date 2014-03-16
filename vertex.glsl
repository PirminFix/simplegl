#version 130

in vec2 position;
in vec2 texcoord;

out vec2 Texcoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 proj;

void main() {
	Texcoord = texcoord;
        gl_Position = model * vec4(position, 0.0, 1.0);
}
