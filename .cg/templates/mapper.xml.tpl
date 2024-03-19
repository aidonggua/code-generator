<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="{{ refs.mapper.Variables.package }}.{{ transformer.Case.Title(transformer.Case.CamelCase(refs.entity.Table)) }}Mapper">
  <resultMap id="BaseResultMap" type="{{ refs.entity.Variables.package }}.{{ transformer.Case.Title(transformer.Case.CamelCase(refs.entity.Table)) }}">
{% for column in table.Columns %}    <id column="{{ column.Name }}" jdbcType="{{ transformer.Type.DbToJDBC(column.Type) }}" property="{{ transformer.Case.CamelCase(column.Name) }}" />
{% endfor %}  </resultMap>
  <sql id="Base_Column_List">
    {% for column in table.Columns %}{{ column.Name }},{% endfor %}
  </sql>
</mapper>