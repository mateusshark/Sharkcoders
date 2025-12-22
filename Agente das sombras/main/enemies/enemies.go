embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"enemy\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/main.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.642665
    y: 0.648588
    z: 0.906667
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 5.0\n"
  "      y: -4.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 22.379894\n"
  "  data: 37.476673\n"
  "  data: 7.966259\n"
  "}\n"
  ""
}
