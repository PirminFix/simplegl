#version 130

in vec3 vert;

void main() {
        gl_Position = vec4(vert, 1);
}
