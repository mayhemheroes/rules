package fuzz

import "strconv"
import "github.com/project-flogo/rules/ruleapi"
import "github.com/project-flogo/rules/config"
import "github.com/project-flogo/rules/common/model"


func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            content := string(bytes)
            ruleapi.GetOrCreateRuleSession(content)
            return 0

        case 1:
            content := string(bytes)
            ruleapi.GetOrCreateRuleSessionFromConfig("mayhem", content)
            return 0

        case 2:
            var test config.RuleDescriptor
            test.UnmarshalJSON(bytes)
            return 0

        case 3:
            content := string(bytes)
            model.RegisterTupleDescriptors(content)
            return 0

        case 4:
            var test model.TupleDescriptor
            test.UnmarshalJSON(bytes)
            return 0

        default:
            content := string(bytes)
            var test model.TupleDescriptor
            test.GetProperty(content)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}