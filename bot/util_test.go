package bot

import "testing"

func Test_trimDescription(t *testing.T) {
	type args struct {
		source string
		limit  int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"html entities", args{"<div class=\"d_post_content_main \"> <div class=\"p_content \"> <cc> <p id=\"post_content_131013371172\" class=\"d_post_content\"> &#x4E0B;&#x4E00;&#x79D2;&#x4E00;&#x9635;&#x5267;&#x70C8;&#x7684;&#x5239;&#x8F66;&#x58F0;&#x6709;&#x4EBA;&#x4ECE;&#x8EAB;&#x540E;&#x4E00;&#x628A;&#x62C9;&#x4F4F;&#x4E86;&#x4ED6;&#x5E03;&#x6EE1;&#x4F24;&#x75D5;&#x7684;&#x540E;&#x80CC;&#x91CD;&#x91CD;&#x843D;&#x5165;&#x4E86;&#x4E00;&#x4E2A;&#x6E29;&#x6696;&#x7684;&#x6000;&#x62B1;</p><br> </cc> <br> </div> <div class=\"core_reply j_lzl_wrapper\"></div></div>", 255}, "下一秒一阵剧烈的刹车声有人从身后一把拉住了他布满伤痕的后背重重落入了一个温暖的怀抱"},
		{"html entities", args{"<p>line1<br>line2<br><br>line3<br><br><br>", 255}, "line1\nline2\nline3"},
	}

	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			if got := trimDescription(item.args.source, item.args.limit); got != item.want {
				t.Errorf("trimDescription() = %v, want: %v", got, item.want)
			}
		})

	}
}
