with col as (
  select column_name as name,
    table_name as tablename,
    table_schema as schemaname,
    udt_name as datatype,
    ordinal_position as ordinalposition,
    coalesce(column_default, '') as defaultvalue,
    coalesce(character_maximum_length, 0) as maxlength,
    coalesce(numeric_precision, 0) as numericprecision,
    coalesce(numeric_scale, 0) as numericscale,
    is_nullable::bool as isnullable,
    is_identity::bool as isidentity,
    case
      when is_generated = 'ALWAYS' then true
      else false
    end as isgenerated,
    '' as comment
  from information_schema.columns
)