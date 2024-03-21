<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="{{.Refs.mapper.Variables.package}}.{{titleCamelCase .Table.Name}}Mapper">
    <resultMap id="BaseResultMap" type="{{.Refs.entity.Variables.package}}.{{titleCamelCase .Table.Name}}">
    {{range .Table.Columns -}}
        {{"    "}}<id column="{{.Name}}" jdbcType="{{dbToJDBC .Type}}" property="{{camelCase .Name}}" />
    {{end -}}
    </resultMap>

    <sql id="Base_Column_List">
    {{"    " -}}
    {{range $i,$v := .Table.Columns -}}
        {{if ne $i 0}},{{end}}{{.Name -}}
    {{end}}
    </sql>
</mapper>