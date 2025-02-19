// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
///////////////////////////////////////////////////////////////////////////////

#include "walkthrough/create_keyset.h"

#include <memory>
#include <string>

#include "gmock/gmock.h"
#include "absl/strings/string_view.h"
#include "tink/aead.h"
#include "tink/aead/aead_config.h"
#include "tink/registry.h"
#include "tink/util/statusor.h"
#include "tink/util/test_matchers.h"

namespace tink_walkthrough {
namespace {

using ::crypto::tink::test::IsOk;
using ::crypto::tink::test::IsOkAndHolds;
using ::crypto::tink::util::StatusOr;
using ::testing::Not;
using ::testing::Test;

class CreateAead256GcmSivKeysetTest : public Test {
 public:
  void TearDown() override { crypto::tink::Registry::Reset(); }
};

TEST_F(CreateAead256GcmSivKeysetTest,
       CreateAead256GcmSivKeysetFailsIfAeadNotRegistered) {
  EXPECT_THAT(CreateAead256GcmSivKeyset(), Not(IsOk()));
}

TEST_F(CreateAead256GcmSivKeysetTest, CreateAead256GcmSivKeysetSucceeds) {
  ASSERT_THAT(crypto::tink::AeadConfig::Register(), IsOk());
  StatusOr<std::unique_ptr<crypto::tink::KeysetHandle>> keyset_handle =
      CreateAead256GcmSivKeyset();
  ASSERT_THAT(keyset_handle, IsOk());
  constexpr absl::string_view plaintext = "Some plaintext";
  constexpr absl::string_view associated_data = "Some associated_data";
  StatusOr<std::unique_ptr<crypto::tink::Aead>> aead =
      (*keyset_handle)->GetPrimitive<crypto::tink::Aead>();
  ASSERT_THAT(aead, IsOk());
  StatusOr<std::string> ciphertext =
      (*aead)->Encrypt(plaintext, associated_data);
  ASSERT_THAT(ciphertext, IsOk());
  EXPECT_THAT((*aead)->Decrypt(*ciphertext, associated_data),
              IsOkAndHolds(plaintext));
}

}  // namespace
}  // namespace tink_walkthrough
