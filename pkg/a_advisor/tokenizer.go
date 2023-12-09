package advisor

import (
	"fmt"
	"vitess.io/vitess/go/vt/sqlparser"
)

// Token basic definition
type Token struct {
	Type int
	Val  string
	i    int
}

// Tokenizer Used to initialize token, different from the Tokenize function,
// this function uses the word segmentation method of vitess
func Tokenizer(sql string) []Token {
	var tokens []Token
	tkn := sqlparser.NewStringTokenizer(sql)
	typ, val := tkn.Scan()
	for typ != 0 {
		if val != nil {
			tokens = append(tokens, Token{Type: typ, Val: string(val)})
		} else {
			if typ > 255 {
				if v, ok := TokenString[typ]; ok {
					tokens = append(tokens, Token{
						Type: typ,
						Val:  v,
					})
				} else {
					tokens = append(tokens, Token{
						Type: typ,
						Val:  "",
					})
				}
			} else {
				tokens = append(tokens, Token{
					Type: typ,
					Val:  fmt.Sprintf("%c", typ),
				})
			}
		}
		typ, val = tkn.Scan()
	}
	return tokens
}

// TokenString sql parser tokens
var TokenString = map[int]string{
	sqlparser.LEX_ERROR:               "",
	sqlparser.UNION:                   "union",
	sqlparser.SELECT:                  "select",
	sqlparser.STREAM:                  "stream",
	sqlparser.INSERT:                  "insert",
	sqlparser.UPDATE:                  "update",
	sqlparser.DELETE:                  "delete",
	sqlparser.FROM:                    "from",
	sqlparser.WHERE:                   "where",
	sqlparser.GROUP:                   "group",
	sqlparser.HAVING:                  "having",
	sqlparser.ORDER:                   "order",
	sqlparser.BY:                      "by",
	sqlparser.LIMIT:                   "limit",
	sqlparser.OFFSET:                  "offset",
	sqlparser.FOR:                     "for",
	sqlparser.ALL:                     "all",
	sqlparser.DISTINCT:                "distinct",
	sqlparser.AS:                      "as",
	sqlparser.EXISTS:                  "exists",
	sqlparser.ASC:                     "asc",
	sqlparser.DESC:                    "desc",
	sqlparser.INTO:                    "into",
	sqlparser.DUPLICATE:               "duplicate",
	sqlparser.KEY:                     "key",
	sqlparser.DEFAULT:                 "default",
	sqlparser.SET:                     "set",
	sqlparser.LOCK:                    "lock",
	sqlparser.KEYS:                    "keys",
	sqlparser.VALUES:                  "values",
	sqlparser.LAST_INSERT_ID:          "last_insert_id",
	sqlparser.NEXT:                    "next",
	sqlparser.VALUE:                   "value",
	sqlparser.SHARE:                   "share",
	sqlparser.MODE:                    "mode",
	sqlparser.SQL_NO_CACHE:            "sql_no_cache",
	sqlparser.SQL_CACHE:               "sql_cache",
	sqlparser.JOIN:                    "join",
	sqlparser.STRAIGHT_JOIN:           "straight_join",
	sqlparser.LEFT:                    "left",
	sqlparser.RIGHT:                   "right",
	sqlparser.INNER:                   "inner",
	sqlparser.OUTER:                   "outer",
	sqlparser.CROSS:                   "cross",
	sqlparser.NATURAL:                 "natural",
	sqlparser.USE:                     "use",
	sqlparser.FORCE:                   "force",
	sqlparser.ON:                      "on",
	sqlparser.USING:                   "using",
	sqlparser.ID:                      "id",
	sqlparser.HEX:                     "hex",
	sqlparser.STRING:                  "string",
	sqlparser.INTEGRAL:                "integral",
	sqlparser.FLOAT:                   "float",
	sqlparser.HEXNUM:                  "hexnum",
	sqlparser.VALUE_ARG:               "?",
	sqlparser.LIST_ARG:                ":",
	sqlparser.COMMENT:                 "",
	sqlparser.COMMENT_KEYWORD:         "comment",
	sqlparser.BIT_LITERAL:             "bit_literal",
	sqlparser.NULL:                    "null",
	sqlparser.TRUE:                    "true",
	sqlparser.FALSE:                   "false",
	sqlparser.OR:                      "||",
	sqlparser.AND:                     "&&",
	sqlparser.NOT:                     "not",
	sqlparser.BETWEEN:                 "between",
	sqlparser.CASE:                    "case",
	sqlparser.WHEN:                    "when",
	sqlparser.THEN:                    "then",
	sqlparser.ELSE:                    "else",
	sqlparser.END:                     "end",
	sqlparser.LE:                      "<",
	sqlparser.GE:                      ">=",
	sqlparser.NE:                      "<>",
	sqlparser.NULL_SAFE_EQUAL:         "<=>",
	sqlparser.IS:                      "is",
	sqlparser.LIKE:                    "like",
	sqlparser.REGEXP:                  "regexp",
	sqlparser.IN:                      "in",
	sqlparser.SHIFT_LEFT:              "<<",
	sqlparser.SHIFT_RIGHT:             ">>",
	sqlparser.DIV:                     "div",
	sqlparser.MOD:                     "mod",
	sqlparser.UNARY:                   "unary",
	sqlparser.COLLATE:                 "collate",
	sqlparser.BINARY:                  "binary",
	sqlparser.UNDERSCORE_BINARY:       "_binary",
	sqlparser.INTERVAL:                "interval",
	sqlparser.JSON_EXTRACT_OP:         "->>",
	sqlparser.JSON_UNQUOTE_EXTRACT_OP: "->",
	sqlparser.CREATE:                  "create",
	sqlparser.ALTER:                   "alter",
	sqlparser.DROP:                    "drop",
	sqlparser.RENAME:                  "rename",
	sqlparser.ANALYZE:                 "analyze",
	sqlparser.ADD:                     "add",
	sqlparser.SCHEMA:                  "schema",
	sqlparser.TABLE:                   "table",
	sqlparser.INDEX:                   "index",
	sqlparser.VIEW:                    "view",
	sqlparser.TO:                      "to",
	sqlparser.IGNORE:                  "ignore",
	sqlparser.IF:                      "if",
	sqlparser.UNIQUE:                  "unique",
	sqlparser.PRIMARY:                 "primary",
	sqlparser.COLUMN:                  "column",
	sqlparser.CONSTRAINT:              "constraint",
	sqlparser.SPATIAL:                 "spatial",
	sqlparser.FULLTEXT:                "fulltext",
	sqlparser.FOREIGN:                 "foreign",
	sqlparser.SHOW:                    "show",
	sqlparser.DESCRIBE:                "describe",
	sqlparser.EXPLAIN:                 "explain",
	sqlparser.DATE:                    "date",
	sqlparser.ESCAPE:                  "escape",
	sqlparser.REPAIR:                  "repair",
	sqlparser.OPTIMIZE:                "optimize",
	sqlparser.TRUNCATE:                "truncate",
	sqlparser.MAXVALUE:                "maxvalue",
	sqlparser.PARTITION:               "partition",
	sqlparser.REORGANIZE:              "reorganize",
	sqlparser.LESS:                    "less",
	sqlparser.THAN:                    "than",
	sqlparser.PROCEDURE:               "procedure",
	sqlparser.TRIGGER:                 "trigger",
	sqlparser.VINDEX:                  "vindex",
	sqlparser.VINDEXES:                "vindexes",
	sqlparser.STATUS:                  "status",
	sqlparser.VARIABLES:               "variables",
	sqlparser.BEGIN:                   "begin",
	sqlparser.START:                   "start",
	sqlparser.TRANSACTION:             "transaction",
	sqlparser.COMMIT:                  "commit",
	sqlparser.ROLLBACK:                "rollback",
	sqlparser.BIT:                     "bit",
	sqlparser.TINYINT:                 "tinyint",
	sqlparser.SMALLINT:                "smallint",
	sqlparser.MEDIUMINT:               "mediumint",
	sqlparser.INT:                     "int",
	sqlparser.INTEGER:                 "integer",
	sqlparser.BIGINT:                  "bigint",
	sqlparser.INTNUM:                  "intnum",
	sqlparser.REAL:                    "real",
	sqlparser.DOUBLE:                  "double",
	sqlparser.FLOAT_TYPE:              "float_type",
	sqlparser.DECIMAL:                 "decimal",
	sqlparser.NUMERIC:                 "numeric",
	sqlparser.TIME:                    "time",
	sqlparser.TIMESTAMP:               "timestamp",
	sqlparser.DATETIME:                "datetime",
	sqlparser.YEAR:                    "year",
	sqlparser.CHAR:                    "char",
	sqlparser.VARCHAR:                 "varchar",
	sqlparser.BOOL:                    "bool",
	sqlparser.CHARACTER:               "character",
	sqlparser.VARBINARY:               "varbinary",
	sqlparser.NCHAR:                   "nchar",
	sqlparser.TEXT:                    "text",
	sqlparser.TINYTEXT:                "tinytext",
	sqlparser.MEDIUMTEXT:              "mediumtext",
	sqlparser.LONGTEXT:                "longtext",
	sqlparser.BLOB:                    "blob",
	sqlparser.TINYBLOB:                "tinyblob",
	sqlparser.MEDIUMBLOB:              "mediumblob",
	sqlparser.LONGBLOB:                "longblob",
	sqlparser.JSON:                    "json",
	sqlparser.ENUM:                    "enum",
	sqlparser.GEOMETRY:                "geometry",
	sqlparser.POINT:                   "point",
	sqlparser.LINESTRING:              "linestring",
	sqlparser.POLYGON:                 "polygon",
	sqlparser.GEOMETRYCOLLECTION:      "geometrycollection",
	sqlparser.MULTIPOINT:              "multipoint",
	sqlparser.MULTILINESTRING:         "multilinestring",
	sqlparser.MULTIPOLYGON:            "multipolygon",
	sqlparser.NULLX:                   "nullx",
	sqlparser.AUTO_INCREMENT:          "auto_increment",
	sqlparser.APPROXNUM:               "approxnum",
	sqlparser.SIGNED:                  "signed",
	sqlparser.UNSIGNED:                "unsigned",
	sqlparser.ZEROFILL:                "zerofill",
	sqlparser.DATABASES:               "databases",
	sqlparser.TABLES:                  "tables",
	sqlparser.NAMES:                   "names",
	sqlparser.CHARSET:                 "charset",
	sqlparser.GLOBAL:                  "global",
	sqlparser.SESSION:                 "session",
	sqlparser.CURRENT_TIMESTAMP:       "current_timestamp",
	sqlparser.DATABASE:                "database",
	sqlparser.CURRENT_DATE:            "current_date",
	sqlparser.CURRENT_TIME:            "current_time",
	sqlparser.LOCALTIME:               "localtime",
	sqlparser.LOCALTIMESTAMP:          "localtimestamp",
	sqlparser.UTC_DATE:                "utc_date",
	sqlparser.UTC_TIME:                "utc_time",
	sqlparser.UTC_TIMESTAMP:           "utc_timestamp",
	sqlparser.REPLACE:                 "replace",
	sqlparser.CONVERT:                 "convert",
	sqlparser.CAST:                    "cast",
	sqlparser.SUBSTR:                  "substr",
	sqlparser.SUBSTRING:               "substring",
	sqlparser.GROUP_CONCAT:            "group_concat",
	sqlparser.SEPARATOR:               "separator",
	sqlparser.VSCHEMA:                 "vschema",
	sqlparser.SEQUENCE:                "sequence",
	sqlparser.MATCH:                   "match",
	sqlparser.AGAINST:                 "against",
	sqlparser.BOOLEAN:                 "boolean",
	sqlparser.LANGUAGE:                "language",
	sqlparser.WITH:                    "with",
	sqlparser.QUERY:                   "query",
	sqlparser.EXPANSION:               "expansion",
	sqlparser.UNUSED:                  "",
}