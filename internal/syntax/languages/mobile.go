package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(SwiftLang)
	syntax.RegisterLanguage(ObjectiveCLang)
	syntax.RegisterLanguage(DartLang)
}

// SwiftLang defines syntax highlighting rules for Swift.
var SwiftLang = &syntax.Language{
	Name:       "Swift",
	Extensions: []string{".swift"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"""[\s\S]*?"""`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(actor|any|as|associatedtype|associativity|async|await|break|case|catch|class|consume|consuming|continue|convenience|default|defer|deinit|didSet|distributed|do|dynamic|else|enum|extension|fallthrough|fileprivate|final|for|func|get|guard|higherThan|if|import|in|indirect|infix|init|inout|internal|is|isolated|lazy|left|let|lowerThan|macro|mutating|nil|nonisolated|nonmutating|none|open|operator|optional|override|postfix|precedence|precedencegroup|prefix|private|protocol|public|repeat|required|rethrows|return|right|safe|self|Self|set|some|static|struct|subscript|super|switch|throw|throws|try|Type|typealias|unowned|unsafe|var|weak|where|while|willSet)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(Any|AnyObject|Array|Bool|Character|Dictionary|Double|Float|Float80|Int|Int8|Int16|Int32|Int64|Never|Optional|Set|String|UInt|UInt8|UInt16|UInt32|UInt64|Void)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|nil)\b`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`@\w+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO][0-7_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:]+`)},
	},
}

// ObjectiveCLang defines syntax highlighting rules for Objective-C.
var ObjectiveCLang = &syntax.Language{
	Name:       "Objective-C",
	Extensions: []string{".m", ".mm"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`@"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(auto|break|case|char|const|continue|default|do|double|else|enum|extern|float|for|goto|if|inline|int|long|register|restrict|return|short|signed|sizeof|static|struct|switch|typedef|union|unsigned|void|volatile|while|_Bool|_Complex|_Imaginary|@interface|@implementation|@protocol|@end|@private|@protected|@public|@package|@try|@catch|@finally|@throw|@synthesize|@dynamic|@property|@selector|@encode|@autoreleasepool|@compatibility_alias|@optional|@required|@class|@defs|self|super|id|Class|SEL|IMP|BOOL|nil|Nil|YES|NO|instancetype|nullable|nonnull|null_unspecified|null_resettable|NS_ASSUME_NONNULL_BEGIN|NS_ASSUME_NONNULL_END|__weak|__strong|__unsafe_unretained|__autoreleasing|__block|__bridge|__bridge_retained|__bridge_transfer)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(NSObject|NSString|NSArray|NSDictionary|NSNumber|NSInteger|NSUInteger|CGFloat|NSMutableArray|NSMutableDictionary|NSMutableString|NSData|NSDate|NSURL|NSError|NSNotification|NSBundle|NSUserDefaults|UIView|UIViewController|UILabel|UIButton|UIImage|UIImageView|UITableView|UICollectionView)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(YES|NO|nil|Nil|NULL|true|false)\b`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`#\s*(import|include|define|undef|ifdef|ifndef|if|else|elif|endif|error|pragma|line)\b.*$`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F]+[uUlL]*\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?[fFlL]?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:@]+`)},
	},
}

// DartLang defines syntax highlighting rules for Dart.
var DartLang = &syntax.Language{
	Name:       "Dart",
	Extensions: []string{".dart"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"""[\s\S]*?"""`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'''[\s\S]*?'''`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(abstract|as|assert|async|await|base|break|case|catch|class|const|continue|covariant|default|deferred|do|dynamic|else|enum|export|extends|extension|external|factory|final|finally|for|Function|get|hide|if|implements|import|in|interface|is|late|library|mixin|new|null|on|operator|part|required|rethrow|return|sealed|set|show|static|super|switch|sync|this|throw|try|typedef|var|void|when|while|with|yield)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(bool|double|dynamic|int|num|Object|String|void|List|Map|Set|Iterable|Future|Stream|Function|Type|Symbol|Null|Never)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null)\b`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`@\w+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:]+`)},
	},
}
