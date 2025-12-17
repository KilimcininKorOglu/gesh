package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(SQLLang)
	syntax.RegisterLanguage(GraphQLLang)
}

// SQLLang defines syntax highlighting rules for SQL.
var SQLLang = &syntax.Language{
	Name:       "SQL",
	Extensions: []string{".sql", ".mysql", ".pgsql", ".sqlite"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`--.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`(?i)\b(SELECT|FROM|WHERE|AND|OR|NOT|IN|IS|NULL|LIKE|BETWEEN|EXISTS|CASE|WHEN|THEN|ELSE|END|AS|ON|JOIN|INNER|LEFT|RIGHT|OUTER|FULL|CROSS|NATURAL|USING|GROUP|BY|HAVING|ORDER|ASC|DESC|LIMIT|OFFSET|UNION|ALL|INTERSECT|EXCEPT|INSERT|INTO|VALUES|UPDATE|SET|DELETE|CREATE|ALTER|DROP|TABLE|INDEX|VIEW|DATABASE|SCHEMA|TRIGGER|PROCEDURE|FUNCTION|IF|BEGIN|COMMIT|ROLLBACK|TRANSACTION|GRANT|REVOKE|PRIMARY|KEY|FOREIGN|REFERENCES|UNIQUE|CHECK|DEFAULT|CONSTRAINT|CASCADE|RESTRICT|AUTO_INCREMENT|SERIAL|IDENTITY|RETURNING|DISTINCT|TOP|FETCH|FIRST|NEXT|ROWS|ONLY|WITH|RECURSIVE|TEMPORARY|TEMP|TRUNCATE|RENAME|ADD|COLUMN|MODIFY|CHANGE)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`(?i)\b(INT|INTEGER|SMALLINT|BIGINT|TINYINT|FLOAT|REAL|DOUBLE|DECIMAL|NUMERIC|CHAR|VARCHAR|TEXT|NCHAR|NVARCHAR|NTEXT|BINARY|VARBINARY|BLOB|CLOB|DATE|TIME|DATETIME|TIMESTAMP|YEAR|BOOLEAN|BOOL|BIT|JSON|JSONB|XML|UUID|ARRAY|ENUM|SET|MONEY|BYTEA|SERIAL|BIGSERIAL|SMALLSERIAL)\b`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`(?i)\b(COUNT|SUM|AVG|MIN|MAX|ROUND|FLOOR|CEIL|CEILING|ABS|POWER|SQRT|MOD|CONCAT|SUBSTRING|SUBSTR|TRIM|LTRIM|RTRIM|UPPER|LOWER|LENGTH|LEN|REPLACE|REVERSE|COALESCE|NULLIF|IFNULL|NVL|CAST|CONVERT|DATE_FORMAT|NOW|CURRENT_DATE|CURRENT_TIME|CURRENT_TIMESTAMP|GETDATE|DATEADD|DATEDIFF|EXTRACT|YEAR|MONTH|DAY|HOUR|MINUTE|SECOND|ROW_NUMBER|RANK|DENSE_RANK|NTILE|LAG|LEAD|FIRST_VALUE|LAST_VALUE|OVER|PARTITION)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`(?i)\b(TRUE|FALSE|NULL)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%<>=!|&^~]+`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`[@:]\w+`)},
	},
}

// GraphQLLang defines syntax highlighting rules for GraphQL.
var GraphQLLang = &syntax.Language{
	Name:       "GraphQL",
	Extensions: []string{".graphql", ".gql"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"""[\s\S]*?"""`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(query|mutation|subscription|fragment|on|type|interface|union|enum|scalar|input|extend|implements|directive|schema|repeatable)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(Int|Float|String|Boolean|ID)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null)\b`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$\w+`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`@\w+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[!:=|&]+`)},
	},
}
