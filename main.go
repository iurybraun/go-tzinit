/*
 * Copyright © 2016-2023 Iury Braun
 * Copyright © 2017-2023 BRAUN TECH
 */

package tzinit

import (
    "os"
)

func init() {
    os.Setenv("TZ", "America/Sao_Paulo")
}
