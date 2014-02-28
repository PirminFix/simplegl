#version 130

in vec2 position;

void main() {
        gl_Position = vec4(vert, 0.0, 1.0);
}
