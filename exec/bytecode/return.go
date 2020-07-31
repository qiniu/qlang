/*
 Copyright 2020 The GoPlus Authors (goplus.org)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package bytecode

const (
	bitsOpReturn        = bitsOp + 3
	bitsOpReturnShift   = bitsInstr - bitsOpReturn
	bitsOpReturnOperand = (1 << bitsOpReturnShift) - 1
)

const (
	bitsRtnNoneOperand  = uint32(0)
	bitsRtnMultiOperand = uint32(1)
	bitsRtnBrkOperand   = uint32(2)
	bitsRtnCtnOperand   = uint32(3)
	bitsRtnGotoOperand  = uint32(4)
)